import { useEintragHandlers } from './shared/useEintragHandlers';
import { EintragTable } from './shared/EintragTable';

export const Schritt2_IdeellerBereich = () => {
  const {
    steuererklaerung,
    errors,
    handleEinnahmeChange,
    handleAusgabeChange,
    addEinnahme,
    addAusgabe,
    removeEinnahme,
    removeAusgabe,
  } = useEintragHandlers('ideellerBereich');



  return (
    <div className="space-y-8">
      <EintragTable
        title="Einnahmen im ideellen Bereich"
        tooltip="Hierzu zählen echte Spenden, Mitgliedsbeiträge und Zuschüsse ohne direkte Gegenleistung."
        eintraege={steuererklaerung.ideellerBereich.einnahmen}
        errors={errors.einnahmen}
        onEintragChange={handleEinnahmeChange}
        onAddEintrag={addEinnahme}
        onRemoveEintrag={removeEinnahme}
        addButtonText="Einnahme hinzufügen"
      />

      <EintragTable
        title="Ausgaben im ideellen Bereich"
        tooltip="Typische Ausgaben sind Verwaltungskosten, Mieten für Vereinsräume oder Kosten für die Mitgliederversammlung."
        eintraege={steuererklaerung.ideellerBereich.ausgaben}
        errors={errors.ausgaben}
        onEintragChange={handleAusgabeChange}
        onAddEintrag={addAusgabe}
        onRemoveEintrag={removeAusgabe}
        addButtonText="Ausgabe hinzufügen"
      />
    </div>
  );
};