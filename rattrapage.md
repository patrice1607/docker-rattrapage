# TP - Rattrapage Docker & CI/CD - Sabry et Patrice

## Objectif du rattrapage

Vous récupérez un dépôt inconnu contenant une API Go.

Sans connaissance préalable de Go, vous devez :

identifier comment l'application démarre ;
construire une image Docker fonctionnelle ;
exécuter l'application dans un conteneur ;
automatiser le build en CI ;
déployer sur Render.

Le but n'est pas de développer l'API mais de la mettre en production.

## Déroulé (1h)

### Étape 1 - Analyse du dépôt (10 min)

Questions à se poser :

Quel langage est utilisé ?
Comment identifier les dépendances ?
Quel est le point d'entrée ?
Quel port est utilisé ?

Indices à rechercher :

go.mod
main.go
Dockerfile éventuel
README

Livrable :

Compléter les questions ci-dessous :
Quel langage ? => 
Point d'entrée du programme ? =>
Port utilisé ? =>
comment installer les dépendances ? =>

### Étape 2 - Dockeriser l'application (20 min)

Objectif :

Créer un Dockerfile permettant de lancer l'API sans installer Go sur la machine.

Contraintes :

utiliser une image officielle ;
récupérer les dépendances ;
compiler l'application ;
démarrer l'API.
Vérification

Quelles commandes taper pour construire et lancer l'image ?

Build :

Run :

### Étape 3 - CI avec GitHub Actions (10 min)

Créer un workflow qui :

construit l'image ;
pousse l'image sur Docker Hub.

### Étape 4 - Déploiement sur Render (10 min)

Objectif :

Déployer directement depuis GitHub avec le webhook Render.

Validation :

L'URL Render répond.
La fournir ici => 

Fournir le lien vers votre dépôt GitHub =>