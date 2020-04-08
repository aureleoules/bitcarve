# BitCarve
BitCarve allows your to carve any type of data on the Bitcoin network **forever**.

## Get started

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

## Requirements
* Go
* bitcoin-daemon and bitcoin-cli (chain data is not required)

## Arguments
### Carve
* --network: main/testnet  
    default=main

* --amount: amount to send to each address  
    default=1 sat

* --fee: Fee sats/byte  
    default=1 sat

* **--utxo**: UTXO ID

* --vout: Position in the UTXO  
    default=0

* **--file**: File to carve

* **--key**: Private key to sign TX

## Decrypt
* **--decrypt**  

* **--txid**: TX storing data  

* --output: data output  
    default=./output

## License

MIT