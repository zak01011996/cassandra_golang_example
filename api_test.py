import requests
import json

s = requests.Session()

# url = "https://wikishop.us/fn/do_auth"
# data = {'email': 'iworld_96@mail.ru', 'password': 'iworldpppoe1996'}
# r = s.post(url, data=json.dumps(data))
# print("Cookies: " + str(s.cookies))
# print("Reg: " + r.text)


# # Get all posts
# url = "http://127.0.0.1:8000"
# r = requests.get(url)
# print("Result: " + r.text)
#
# Registration
url = "http://127.0.0.1:8000/signup"
data = {'login': 'test', 'pass': 'test', 'name': 'Test'}
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
r = s.post(url, data=json.dumps(data), headers=headers)
print("Result: " + r.text)
#
# # Authorization
# url = "http://127.0.0.1:8000/signin"
# data = {'login': 'test', 'pass': 'test'}
# r = s.post(url, data=json.dumps(data))
# print("Auth: " + r.text)

# # Logout
# url = "http://127.0.0.1:8000/signout"
# headers = {'auth_token': 'b8928e68-111b-11e5-90e1-30f9edaddb1c'}
# r = s.get(url, headers=headers)
# print("Auth: " + r.text)
#
# # Create Post
# url = "http://127.0.0.1:8000/posts"
# headers = {'auth_token': '348ec76f-111e-11e5-957d-30f9edaddb1c'}
# data = {'title': 'test_2', 'msg': 'testtestesttestesttest_22222'}
# r = s.post(url, headers=headers, data=json.dumps(data))
# print("Result: " + r.text, )

# # Delete Post
# url = "http://127.0.0.1:8000/posts/40f1c83b-111e-11e5-957e-30f9edaddb1c"
# headers = {'auth_token': '92dfbb66-111c-11e5-98b6-30f9edaddb1c'}
# r = s.delete(url, headers=headers)
# print("Result: " + r.text, )
