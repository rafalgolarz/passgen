# Passgen

**API to generate secure passwords**

## Configurarion

**config.toml** stores minimum default required settings for generated passwords

Currently it contains two sections: 

```
[default]
min_length = 8
min_special_characters = 2
min_digits = 2
min_lowercase = 1
min_uppercase = 1
results = 1

[strong]
min_length = 16
min_special_characters = 4
min_digits = 4
min_lowercase = 2
min_uppercase = 2
results = 1
```

You can choose which section should be loaded as default one, in **definitions.go**:

```
passwordType = "default"
```

Each of the params can be overwritten from the url.

## Running

```
docker build -f Dockerfile-dev -t rafalgolarz/passgen . 
docker run --rm -p 8080:8080 --name=passgen rafalgolarz/passgen
```
Dockerfile-dev has debug flags on. Use Dockerfile to have them off or set them with -e param:

```
docker run --rm -p 8080:8080 --name=passgen -e GIN_MODE=release rafalgolarz/passgen
```

Next, open the url: 

http://localhost:8080/v1/passwords

By default, it generates one password meeting criteria defined in config.toml but you can overwrite any of the params.

Generate 3 passwords:

http://localhost:8080/v1/passwords?res=3

Generate 20 passwords. Each of the passwords should have:
- minimum length of 25 characters
- minimum of 2 special characters
- minimum of 2 digits
- minimum 4 lower and 4 upper case letters

http://localhost:8080/v1/passwords/?min-length=25&min-specials=2&min-digits=2&min-lowers=4&min-uppers=4&res=20
