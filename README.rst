====
Loch
====

There's maybe something hidden in the loch.

.. image:: misc/loch.jpg
   :scale: 25 %
   :alt: Loch
   :align: center

Loch is a simple utility to encrypt and decrypt files
using NaCL's secretbox


Usage
=====

Encrypting a file:

::
    
    $ loch --out source_file.encrypted encrypt source_file


Decrypting a file:

::
    
    $ loch --out source_file.decrypted decrypt source_file.decrypted


You can provide a text file containing the secret (instead of typing it out
in the console) by using the `key` flag.
