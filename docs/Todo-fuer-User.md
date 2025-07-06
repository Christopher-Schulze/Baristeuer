# Externe Binärdaten bereitstellen

Die Repository-Richtlinien sehen vor, dass große oder binäre Dateien (z. B. Beispiel-PDFs) nicht direkt eingecheckt werden. Sollten für eigene Tests oder Dokumentation Beispieldateien benötigt werden, können sie außerhalb des Repositories bereitgestellt werden.

1. Die Dateien zunächst lokal erzeugen oder sammeln.
2. Anschließend einen Download-Link (z. B. in einem privaten Cloud-Speicher) anlegen.
3. Den Link in der Projektbeschreibung oder im Pull Request erwähnen, damit andere darauf zugreifen können.
4. Große Dateien sollten nach Möglichkeit komprimiert (ZIP) und mit einer kurzen Erläuterung versehen werden.

So bleiben die Versionskontrolle und der Repository-Umfang schlank, während trotzdem alle notwendigen Daten verfügbar sind.

## Offizielle Formulare einbinden

Im Verzeichnis `internal/pdf/templates` befinden sich lediglich Platzhalterdateien.
Laden Sie daher die amtlichen PDF-Formulare von der Finanzverwaltung oder aus ELSTER
herunter und speichern Sie sie **lokal** unter dem gleichen Dateinamen in diesem
Ordner. Die Dateinamen müssen exakt den Platzhaltern entsprechen (z. B. `kst1.pdf`).

Gehen Sie dazu wie folgt vor:
1. Öffnen Sie das [ELSTER-Portal](https://www.elster.de/eportal/formulare-leistungen/alleformulare) oder die Website Ihrer
   Landesfinanzverwaltung.
2. Suchen Sie nach den Formularen **KSt 1**, **KSt 1F**, **Anlage Gem**, **Anlage GK** und **Anlage Sport**.
3. Laden Sie jede PDF-Datei herunter und benennen Sie sie genau
   `kst1.pdf`, `kst1f.pdf`, `anlage_gem.pdf`, `anlage_gk.pdf` und `anlage_sport.pdf`.
4. Speichern Sie diese Dateien anschließend im Ordner `internal/pdf/templates/`.

Beim Erzeugen der Formulare greift der Generator automatisch auf diese Vorlagen
zurück, sofern sie vorhanden sind.

**Hinweis:** Aus Lizenzgründen dürfen die Original-PDFs nicht ins Repository
eingecheckt werden. Legen Sie die Dateien `kst1.pdf`, `kst1f.pdf`,
`anlage_gem.pdf`, `anlage_gk.pdf` und `anlage_sport.pdf` daher manuell in
`internal/pdf/templates/` ab und schließen Sie sie von einem Commit aus.
