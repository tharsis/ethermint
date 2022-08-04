import pytest

from .utils import (
    KEYS,
    CONTRACTS,
    ADDRS,
    send_transaction,
    deploy_contract,
    w3_wait_for_new_blocks,
)

txhash = ""

def test_block(ethermint, geth):
    get_blocks(ethermint, geth, False)
    get_blocks(ethermint, geth, True)

def get_blocks(ethermint, geth, with_transactions):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockByNumber", ['0x0', with_transactions])
    
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockByNumber", ['0x2710', with_transactions])
    
    ethermint_blk = ethermint.w3.eth.get_block(1)
    # Get existing block, no transactions
    eth_rsp = eth_rpc.make_request(
        "eth_getBlockByHash", [ethermint_blk['hash'].hex(), with_transactions]
    )
    geth_rsp = geth_rpc.make_request(
        "eth_getBlockByHash", ["0x124d099a1f435d3a6155e5d157ff1078eaefb742435892677ee5b3cb5e6fa055", with_transactions]
    )
    res, err = same_types(eth_rsp, geth_rsp)
    assert res, err

    # Get not existing block
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockByHash", ['0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd', with_transactions])

    # Bad call
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockByHash", ['0', with_transactions])


def test_accounts(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_accounts", [])

def test_syncing(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_syncing", [])

def test_coinbase(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_coinbase", [])

def test_max_priority_fee(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_maxPriorityFeePerGas", [])


def test_gas_price(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_gasPrice", [])


def test_block_number(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_blockNumber", [])

def test_balance(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", 'latest'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0", 'latest'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x9907a0cf64ec9fbf6ed8fd4971090de88222a9ac", 'latest'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", '0x0'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0",'0x0'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x9907a0cf64ec9fbf6ed8fd4971090de88222a9ac",'0x0'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", '0x10000'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0",'0x10000'])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBalance", ["0x9907a0cf64ec9fbf6ed8fd4971090de88222a9ac",'0x10000'])

def deploy_and_wait(w3, number=1):
    contract = deploy_contract(
        w3,
        CONTRACTS["TestERC20A"],
    )

    w3_wait_for_new_blocks(w3, number)
    return contract

def test_getStorageAt(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getStorageAt", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", '0x0', 'latest'])

    contract = deploy_and_wait(ethermint.w3)
    res = eth_rpc.make_request(
        'eth_getStorageAt', [contract.address, '0x0', 'latest']
    )
    expected = '0x00000000000000000000000000000000000000000000000000120a0b063499d4'
    res, err = same_types(res['result'], expected)
    assert res, err

def send_and_get_hash(w3, tx_value=10):
    # Do an ethereum transfer
    gas_price = w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": tx_value, "gasPrice": gas_price}
    return send_transaction(w3, tx, KEYS["validator"])["transactionHash"].hex()

def test_getProof(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getProof", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", ['0x0'], 'latest'])

    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getProof", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", ['0x0'], '0x32'])

    _ = send_and_get_hash(ethermint.w3)

    proof = eth_rpc.make_request('eth_getProof', [ADDRS["validator"], ['0x0'], 'latest'])
    expected = {
    'address': '0x4CB06C43fcdABeA22541fcF1F856A6a296448B6c',
    'accountProof': ['0xf90211a03841a7ddd65c70c94b8efa79190d00f0ab134b26f18dcad508f60a7e74559d0ba0464b07429a05039e22931492d6c6251a860c018ea390045d596b1ac11b5c7aa7a011f4b89823a03c9c4b5a8ab079ee1bc0e2a83a508bb7a5dc7d7fb4f2e95d3186a0b5f7c51c3b2d51d97f171d2b38a4df1a7c0acc5eb0de46beeff4d07f5ed20e19a0b591a2ce02367eda31cf2d16eca7c27fd44dbf0864b64ea8259ad36696eb2a04a02b646a7552b8392ae94263757f699a27d6e9176b4c06b9fc0a722f893b964795a02df05d68bceb88eebf68aafde61d10ab942097afc1c58b8435ffd3895358a742a0c2f16143c4d1db03276c433696dddb3e9f3b113bcd854b127962262e98f43147a0828820316cc02bfefd899aba41340659fd06df1e0a0796287ec2a4110239f6d2a050496598670b04df7bbff3718887fa36437d6d8c7afb4eff86f76c5c7097dcc4a0c14e9060c6b3784e35b9e6ae2ad2984142a75910ccc89eb89dc1e2f44b6c58c2a009804db571d0ce07913e1cbacc4f1dc4fb8265c936f5c612e3a47e91c64d8e9fa063d96f38b3cb51b1665c6641e25ffe24803f2941e5df79942f6a53b7169647e4a0899f71abb18c6c956118bf567fac629b75f7e9526873e429d3d8abb6dbb58021a00fd717235298742623c0b3cafb3e4bd86c0b5ab1f71097b4dd19f3d6925d758da0096437146c16097f2ccc1d3e910d65a4132803baee2249e72c8bf0bcaaeb37e580',
                     '0xf90151a097b17a89fd2c03ee98cb6459c08f51b269da5cee46650e84470f62bf83b43efe80a03b269d284a4c3cf8f8deacafb637c6d77f607eec8d75e8548d778e629612310480a01403217a7f1416830c870087c524dabade3985271f6f369a12b010883c71927aa0f592ac54c879817389663be677166f5022943e2fe1b52617a1d15c2f353f27dda0ac8d015a9e668f5877fcc391fae33981c00577096f0455b42df4f8e8089ece24a003ba34a13e2f2fb4bf7096540b42d4955c5269875b9cf0f7b87632585d44c9a580a0b179e3230b07db294473ae57f0170262798f8c551c755b5665ace1215cee10ca80a0552d24252639a6ae775aa1df700ffb92c2411daea7286f158d44081c8172d072a0772a87d08cf38c4c68bfde770968571abd16fd3835cb902486bd2e515d53c12d80a0413774f3d900d2d2be7a3ad999ffa859a471dc03a74fb9a6d8275455f5496a548080',
                     '0xf869a020d13b52a61d3c1325ce3626a51418adebd6323d4840f1bdd93906359d11c933b846f8440180a01ab7c0b0a2a4bbb5a1495da8c142150891fc64e0c321e1feb70bd5f881951f7ea0551332d96d085185ab4019ad8bcf89c45321e136c261eb6271e574a2edf1461f'
                     ],
    'balance': 0,
    'codeHash': '0x551332d96d085185ab4019ad8bcf89c45321e136c261eb6271e574a2edf1461f',
    'nonce': 1,
    'storageHash': '0x1ab7c0b0a2a4bbb5a1495da8c142150891fc64e0c321e1feb70bd5f881951f7e',
    'storageProof': [{
            'key': '0x00',
            'value': '0x48656c6c6f00000000000000000000000000000000000000000000000000000a',
            'proof': ['0xf9019180a01ace80e7bed79fbadbe390876bd1a7d9770edf9462049ef8f4b555d05715d53ea049347a3c2eac6525a3fd7e3454dab19d73b4adeb9aa27d29493b9843f3f88814a085079b4abcd07fd4a5d6c52d35f4c4574aecc85830e90c478ca8c18fcbe590de80a02e3f8ad7ea29e784007f51852b9c3e470aef06b11bac32586a8b691134e4c27da064d2157a14bc31f195f73296ea4dcdbe7698edbf3ca81c44bf7730179d98d94ca09e7dc2597c9b7f72ddf84d7eebb0fe2a2fa2ab54fe668cd14fee44d9b40b1a53a0aa5d4acc7ac636d16bc9655556770bc325e1901fb62dc53770ef9110009e080380a0d5fde962bd2fb5326ddc7a9ca7fe0ee47c5bb3227f838b6d73d3299c22457596a08691410eff46b88f929ef649ea25025f62a5362ca8dc8876e5e1f4fc8e79256d80a0673e88d3a8a4616f676793096b5ae87cff931bd20fb8dd466f97809a1126aad8a08b774a45c2273553e2daf4bbc3a8d44fb542ea29b6f125098f79a4d211b3309ca02fed3139c1791269acb9365eddece93e743900eba6b42a6a8614747752ba268f80',
                      '0xf891808080a0c7d094301e0c54da37b696d85f72de5520b224ab2cf4f045d8db1a3374caf0488080a0fc5581783bfe27fab9423602e1914d719fd71433e9d7dd63c95fe7e58d10c9c38080a0c64f346fc7a21f6679cba8abdf37ca2de8c4fcd8f8bcaedb261b5f77627c93908080808080a0ddef2936a67a3ac7d3d4ff15a935a45f2cc4976c8f0310aed85daf763780e2b480',
                      '0xf843a0200decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563a1a048656c6c6f00000000000000000000000000000000000000000000000000000a'
                      ]
        }
    ]
    }
    res, err = same_types(proof['result'], expected)
    assert res, err

def test_getCode(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getCode", ["0x57f96e6b86cdefdb3d412547816a82e3e0ebf9d2", "latest"])

    # Do an ethereum transfer
    contract = deploy_and_wait(ethermint.w3)
    code = eth_rpc.make_request(
        'eth_getCode', [contract.address, "latest"]
    )
    expected = {
    "id": "4",
    "jsonrpc": "2.0",
    "result": "0x"
    }    
    res, err = same_types(code, expected)
    assert res, err


def test_getBlockTransactionCountByNumber(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockTransactionCountByNumber", ["0x0"])

    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getBlockTransactionCountByNumber", ["0x100"])

    tx_hash = send_and_get_hash(ethermint.w3)

    tx_res = eth_rpc.make_request('eth_getTransactionByHash', [tx_hash])
    block_hash = tx_res['result']['blockNumber']
    block_res = eth_rpc.make_request('eth_getBlockTransactionCountByNumber', [block_hash])
    

    expected = {    
    "id": "1",
    "jsonrpc": "2.0",
    "result": "0x0"
    }
    res, err = same_types(block_res, expected)
    assert res, err


# TODO: BYHASH
# def test_getBlockTransactionCountByHash


def test_getTransaction(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getTransactionByHash", ["0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"])

    tx_hash = send_and_get_hash(ethermint.w3)

    tx_res = eth_rpc.make_request('eth_getTransactionByHash', [tx_hash])

    expected = {
    "jsonrpc": "2.0",
    "id": 0,
    "result": {
        "hash": "0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b",
        "blockHash": "0x1d59ff54b1eb26b013ce3cb5fc9dab3705b415a67127a003c3e61eb445bb8df2",
        "blockNumber": "0x5daf3b",
        "from": "0xa7d9ddbe1f17865597fbd27ec712455208b6b76d",
        "gas": "0xc350",
        "gasPrice": "0x4a817c800",
        "input": "0x68656c6c6f21",
        "nonce": "0x15",
        "r": "0x1b5e176d927f8e9ab405058b2d2457392da3e20f328b16ddabcebc33eaac5fea",
        "s": "0x4ba69724e8f69de52f0125ad8b3c5c2cef33019bac3249e2c0a2192766d1721c",
        "to": "0xf02c1c8e6114b1dbe8937a39260b5b0a374432bb",
        "transactionIndex": "0x41",
        "v": "0x25",
        "value": "0xf3dbb76162000"
    }
    }
    res, err = same_types(tx_res, expected)
    assert res, err

# NOT IMPLEMENTED
# def test_getTransactionRaw(ethermint, geth):
#     eth_rpc = ethermint.w3.provider
#     geth_rpc = geth.w3.provider
#     eth_rsp, geth_rsp = make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getRawTransactionByHash", ["0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"])
#     res, err = same_types(eth_rsp, geth_rsp)
#     print(eth_rsp)
#     print(geth_rsp)
#     assert res, err

def test_getTransactionReceipt(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_getTransactionReceipt", ["0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060"])

    tx_hash = send_and_get_hash(ethermint.w3)

    tx_res = eth_rpc.make_request('eth_getTransactionReceipt', [tx_hash])
    expected = {
    'blockHash': '0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd',
    'blockNumber': 46147,
    'contractAddress': None,
    'cumulativeGasUsed': 21000,
    'from': '0xA1E4380A3B1f749673E270229993eE55F35663b4',
    'gasUsed': 21000,
    'logs': [],
    'logsBloom': '0x000000000000000000000000000000000000000000000000...0000',
    'status': 1, # 0 or 1
    'to': '0x5DF9B87991262F6BA471F09758CDE1c0FC1De734',
    'transactionHash': '0x5c504ed432cb51138bcf09aa5e8a410dd4a1e204ef84bfed1be16dfba1b22060',
    'transactionIndex': 0,
    }
    res, err = same_types(tx_res['result'], expected)
    assert res, err

def test_feeHistory(ethermint, geth):
    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_feeHistory", [4, 'latest',[10,90]])

    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_feeHistory", [4, '0x5000',[10,90]])

    fee_history = eth_rpc.make_request('eth_feeHistory', [4, 'latest',[10,90]])
    expected = {
        'oldestBlock': 3,
        'reward': [[220, 7145389], [1000000, 6000213], [550, 550], [125, 12345678]],
        'baseFeePerGas': [202583058, 177634473, 155594425, 136217133, 119442408],
        'gasUsedRatio': [0.007390479689642084, 0.0036988514889990873, 0.0018512333048507866, 0.00741217041320997]
    }
    res, err = same_types(fee_history['result'], expected)
    assert res, err

def test_estimateGas(ethermint, geth):
    tx = {"to": ADDRS["community"], "from": ADDRS["validator"]}

    eth_rpc = ethermint.w3.provider
    geth_rpc = geth.w3.provider
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_estimateGas", [tx])
    make_same_rpc_calls(eth_rpc, geth_rpc, "eth_estimateGas", [{}])

def make_same_rpc_calls(rpc1, rpc2, method, params):
    res1 = rpc1.make_request(
        method, params
    )
    res2 = rpc2.make_request(
        method, params
    )
    res, err = same_types(res1, res2)
    assert res, err

def same_types(objectA, objectB):

    if isinstance(objectA, dict):
        if not isinstance(objectB, dict):
            return False, 'A is dict, B is not'
        keys = list(set(list(objectA.keys()) +  list(objectB.keys())))
        for key in keys:
            if key in objectB and key in objectA:
                if not same_types(objectA[key], objectB[key]):
                    return False, key + ' key on dict are not of same type'
            else:
                return False, key + ' key on json is not in both results'
        return True, ""
    elif isinstance(objectA, list):
        if not isinstance(objectB, list):
            return False, 'A is list, B is not'
        if len(objectA) == 0 and len(objectB) == 0:
            return True, ""
        if len(objectA) > 0 and len(objectB) > 0:
            return same_types(objectA[0], objectB[0])
        else:
            return True
    elif objectA is None and objectB is None:
        return True, ""
    elif type(objectA) is type(objectB):
        return True, ""
    else:
        return False, 'different types. A is type '+ type(objectA).__name__ + ' B is type '+ type(objectB).__name__