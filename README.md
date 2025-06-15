# EDUVA - Core

> Service core de l'application EDUVA  

## Requis

- docker / docker compose installés  
- git clone du projet  
- la commande `make` doit fonctionner
- traefik doit être installé ([repo](https://github.com/nsevendev/infra-traefik)) et demarré.
- creer les fichier `.env` et `.env.test` à la racine du projet en copiant  
le contenu des fichiers `.env.dist` pour le '.env' et le `.env.test.dist` pour le '.env.test'

## Installation dev

```bash
# lancer les containers
$ make up 

# lancer les containers et (re) build l'image
$ make upb

# arrêter les containers
$ make down
```
- Pour voir les autres commandes disponibles, exécuter `make` ou `make help` dans le terminal

## Swagger  

- Pour voir la documentation sur votre navigateur, afficher les logs de l'appication  
et vous verrez l'url du swagger, cliquez dessus pour l'ouvrir.
