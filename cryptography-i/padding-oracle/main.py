import code
import requests

from utils import guess_helper

URL = 'http://crypto-class.appspot.com/po'
ORIGINAL_CIPHERTEXT = 'f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4'

bytes = bytearray.fromhex(ORIGINAL_CIPHERTEXT)

# code.interact(local=dict(globals(), **locals()))

if False:
    r = requests.get(URL, params={'er': ORIGINAL_CIPHERTEXT})
    code.interact(local=dict(globals(), **locals()))
else:
    cipher = ORIGINAL_CIPHERTEXT
    plain = []
    for i in range(0, len(ORIGINAL_CIPHERTEXT) - 16):
        solved = False
        for guess in range(0, 256):
            cipher = guess_helper(cipher, i, guess)
            r = requests.get(URL, params={'er': ORIGINAL_CIPHERTEXT})
            code.interact(local=dict(globals(), **locals()))
            if r.status_code == 404:
                solved = True
                plain << chr(guess)
                break
        if not solved:
            raise(f'No guess was right in index {i}')
    print(''.join(plain.reverse))
