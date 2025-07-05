# Externe Binärdaten bereitstellen

Die Repository-Richtlinien sehen vor, dass große oder binäre Dateien (z. B. Beispiel-PDFs) nicht direkt eingecheckt werden. Sollten für eigene Tests oder Dokumentation Beispieldateien benötigt werden, können sie außerhalb des Repositories bereitgestellt werden.

1. Die Dateien zunächst lokal erzeugen oder sammeln.
2. Anschließend einen Download-Link (z. B. in einem privaten Cloud-Speicher) anlegen.
3. Den Link in der Projektbeschreibung oder im Pull Request erwähnen, damit andere darauf zugreifen können.
4. Große Dateien sollten nach Möglichkeit komprimiert (ZIP) und mit einer kurzen Erläuterung versehen werden.

So bleiben die Versionskontrolle und der Repository-Umfang schlank, während trotzdem alle notwendigen Daten verfügbar sind.

## Offizielle Formulare einbinden

Im Verzeichnis `internal/pdf/templates` befinden sich nur Platzhalterdateien. Um die amtlichen PDF-Formulare zu verwenden, kopieren Sie die Originale mit dem gleichen Dateinamen (z. B. `kst1.pdf`) in dieses Verzeichnis. Beim Erzeugen der Formulare greift der Generator automatisch auf diese Vorlagen zurück.
