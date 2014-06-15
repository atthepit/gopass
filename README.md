gopass
======

Gopass is a simple and easy to use open-source password administrator for Unix and Linux!

How to
-----------

Gopas it's really easy to use!

Create a user:
  >gopass -u username

Generate and store a secure password for a site:
  >gopass -n sitename -u username

Recover a user's site password:
  >gopass -s sitename -u username

Reset a user's site password:
  >gopass -r sitename -u username

Remove user's site password:
  >gopass -d sitename -u username
  
List all the user's sites stored:
  >gopass -l username

Destroy all user information:
  >gopass -d username
