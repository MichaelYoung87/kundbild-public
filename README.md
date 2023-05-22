# Projekt Kundbild
## Nödvändiga förutsättningar för att kunna testköra applikationen:
```go
1. Version 1.20+ - Go programspråk - https://go.dev/dl
2. Valfri programutvecklingsmiljö - Jag använde mig av Visual Studio Code - https://code.visualstudio.com/download
3. Docker/Docker Engine - För att mysql databas/jenkins containers genom docker-compose.yml ska kunna skapas 
- https://docs.docker.com/get-docker/ - https://docs.docker.com/engine/install/
4. Node.js och npm - Senaste versionen - Krävs för att kunna uppdatera/köra React frontend som ligger i ./web/
5. Valfritt databasprogram för att kunna connecta till mysql databasen och för att se vad som finns sparat, jag använder 
själv DBeaver Community - https://dbeaver.io/download/
6. Konfigurationsdetaljerna för att connecta till mysql databasen hittar ni under # Database Connection Details inuti 
.env filen som befinner sig i root av projektet.
```
## Instruktioner för installation:
```go
1. Navigera med terminal till en valfri folder där ni vill spara projektet.
2. Skriv in 'git clone https://github.com/MichaelYoung87/kundbild-public.git' i terminal och tryck enter.
3. Navigera med terminal till det nedhämtade projektets root folder, 'cd kundbild-public' i terminal och tryck enter.
4. Skriv in 'go mod download' i terminal och tryck enter.
5. Skriv in 'go mod tidy' i terminal och tryck enter.
6. Om ni kör med Windows se till så att Docker Engine är igång.
7. Om ni fortfarande befinner er i projektets root folder, skriv in 'docker-compose up -d' i terminal och tryck enter 
('sudo docker-compose up -d' för Linux).
8. Navigera med terminal till det nedhämtade projektets web folder, skriv in 'cd web' i terminal och tryck enter.
9. Skriv in 'npm install' i terminal och tryck enter. 
```
## Instruktioner för körning av backend/frontend:
```go
1. Gå in på Extensions och sök på 'Go' i sökfältet och installera.
2. Markera main.go som ligger i foldern cmd, tryck på 'Run Code' ikonen uppe till höger i Visual Studio Code (Ctrl + Alt + N).
3. Navigera med terminal till det nedhämtade projektets web folder igen om ni inte fortfarande är där, skriv in 'npm start' 
och tryck enter.
4. Om inte frontend med React automatiskt öppnar upp webbläsaren när React har startats upp kan ni skriva in 
'http://localhost:3000/' manuellt i adressfältet.
```
## Övriga kommandon till docker-compose:
```go
1. För att starta: Skriv in 'docker-compose up -d' i terminal och tryck enter ('sudo docker-compose up -d' för Linux).
2. För att kolla loggarna: Skriv in 'docker-compose logs -f' i terminal och tryck enter ('sudo docker-compose logs -f' för Linux).
3. För att stänga ner: Skriv in 'docker-compose down' i terminal och tryck enter ('sudo docker-compose down' för Linux).
4. För att stänga ner och rensa volymerna (sparad data): Skriv in 'docker-compose down -v' i terminal och tryck enter 
('sudo docker-compose down -v' för Linux).
```
## Kort översikt om hur applikationen fungerar:
```go
1. Skriv in ett nummer mellan 1-83 i fältet för People samt 1-60 i fältet för Planets.
2. Tryck på 'Fetch' knappen för att se information som befinner sig på dessa API:er.
3. Om de två API:erna är hårdkodade att matcha i application\services\flagged_customers_service.go i funktionen 'CheckMatch' 
så kan man trycka på 'Save Flag' knappen så sparas det i sql databasen.
4. 'Save Flag' aktiveras endast om informationen om de två olika API:erna finns hårdkodade i backend.
5. Man kan även trycka på 'Save Linked' knappen oavsett om de matchar i hårdkodningen eller inte.
6. Meddelande att informationen har sparats i databasen visas om man lagt till en ny flaggning eller linkning i databasen.
7. Ni kan rensa bort nummerna ni skrev in i fälten People och Planets för att rensa informationen på skärmen.
```
## Stänga ner backend/frontend:
```go
1. Backend - Tryck på 'Stop Code Run' ikonen uppe till höger i Visual Studio Code (Ctrl + Alt + M).
2. Frontend React - Gå till terminalfönstret där React körs, brukar stå med grön text 'Compiled successfully!' och övrig 
information. Tryck in 'Ctrl + C' i terminal för att stänga av React.
```