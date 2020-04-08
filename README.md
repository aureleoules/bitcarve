# BitCarve
BitCarve allows your to carve any type of data on the Bitcoin network **forever**.

## Get started

**Disclaimer**: To avoid data corruption, there can be no change of the UTXO. You will lose the whole amount in the UTXO.  

To compute the minimum required amount of satoshis in the UTXO:  
`amount = 2 * fileSize (bytes) + fileSize (bytes) % 25`


```bash
$ go build

# Carve data
$ ./bitcarve \
    --file image.jpg  \
    --network testnet \ 
    --amount 1 \
    --fee 1 \
    --utxo your_utxo_id \
    --key your_private_key

Successfully carved data on the Bitcoin network.
TxID: ea33b22a9f63581da342999f3f47ef2fe886e7ddc351fea2ac3168b0fc552ede

# Decrypt data
$ ./bitcarve \
    --decrypt \
    --txid ea33b22a9f63581da342999f3f47ef2fe886e7ddc351fea2ac3168b0fc552ede \
    --network testnet \
    --output myimage.jpg

Retrieved data.
```

### Docker

```bash
$ docker build --tag bitcarve .
$ docker run -it bitcarve \
    --file image.jpg  \
    --network testnet \ 
    --amount 1 \
    --fee 1 \
    --utxo your_utxo_id \
    --key your_private_key

Successfully carved data on the Bitcoin network.
TxID: ea33b22a9f63581da342999f3f47ef2fe886e7ddc351fea2ac3168b0fc552ede
```

## Requirements
* Go
* bitcoin-daemon and bitcoin-cli (chain data is not required)

## Arguments
### Carve
* **--utxo**: UTXO ID  

* **--file**: File to carve

* **--key**: Private key to sign TX

* --network: main/testnet  
    default=main

* --amount: amount to send to each address  
    default=1 sat

* --fee: Fee sats/byte  
    default=1 sat


* --vout: Position in the UTXO  
    default=0


## Decrypt
* **--decrypt**  

* **--txid**: TX storing data  

* --output: data output  
    default=./output

## How it works

To store large files (kilobytes) on the Bitcoin network, the `OP_RETURN` cannot be used because of its 40 bytes limit.

Instead, `bitcarve` stores data in tx output addresses. Each address can store 20 bytes of data, 1 version byte, and 4 bytes of checksum.   

Your file is split into chunks of 20 bytes, hashed to compute a checksum, versionned and converted to base58 in order to produce valid Bitcoin addresses.

Each address is added to the bitcoin transaction with 1 sat (0.00000001 BTC).

The transaction can be signed and broadcast to the network.

Your data is now stored forever on Bitcoin.

## License

MIT