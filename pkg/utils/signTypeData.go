package utils

const LoginTypedData = `
{
    "types": {
	  "EIP712Domain": [
				{
					"name": "name",
					"type": "string"
				},
				{
					"name": "chainId",
					"type": "uint256"
				}
			],
	
        "Message": [
			{
                "name": "type",
                "type": "string"
            },
            {
                "name": "timestamp",
                "type": "uint256"
            },
            {
                "name": "address",
                "type": "address"
            },
			{
                "name": "nonce",
                "type": "string"
            }
        ]
    },
    "domain": {
        "name": "mm",
        "chainId": "%d"
    },
    "primaryType": "Message",
    "message": {
        "type": "LOGIN",
        "timestamp": "%d",
        "address": "%s",
        "nonce": "%s"
    }
}
`
