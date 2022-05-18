# totpg

A simple and small cli tool for generating a one time Password using the Base32 encoded secret key associated with the Time-based One Time Password. This tool was written using the (One Time Password utilities Go)[https://github.com/pquerna/otp].

# Why

This utility was written to provide a simple way to generate One Time Passwords. The intent is to use these inside of continuous integration environments where 2 factor authentication is required. 

# Example

```shell
totpg -secret MYSECRETVALUE
111111
```