import { useEintragHandlers } from './shared/useEintragHandlers';
import { EintragTable } from './shared/EintragTable';

export const Schritt3_Vermoegensverwaltung = () => {
  const {
    steuererklaerung,
    errors,
    handleEinnahmeChange,
    handleAusgabeChange,
    addEinnahme,
    addAusgabe,
    removeEinnahme,
    removeAusgabe,
  } = useEintragHandlers('vermoegensverwaltung');

  return (
    <div className="space-y-8">
      <EintragTable
        title="Einnahmen in der Vermögensverwaltung"
        tooltip="Einnahmen aus Kapitalvermögen (z.B. Zinsen, Dividenden) oder langfristiger Vermietung & Verpachtung."
        eintraege={steuererklaerung.vermoegensverwaltung.einnahmen}
        errors={errors.einnahmen}
        onEintragChange={handleEinnahmeChange}
        onAddEintrag={addEinnahme}
        onRemoveEintrag={removeEinnahme}
        addButtonText="Einnahme hinzufügen"
      />

      <EintragTable
        title="Ausgaben in der Vermögensverwaltung"
        tooltip="Kosten, die direkt mit den Einnahmen der Vermögensverwaltung zusammenhängen (z.B. Kontoführungsgebühren, Reparaturen an vermieteten Objekten)."
        eintraege={steuererklaerung.vermoegensverwaltung.ausgaben}
        errors={errors.ausgaben}
        onEintragChange={handleAusgabeChange}
        onAddEintrag={addAusgabe}
        onRemoveEintrag={removeAusgabe}
        addButtonText="Ausgabe hinzufügen"
      />
    </div>
  );
};