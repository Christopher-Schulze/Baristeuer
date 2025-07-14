import { useEintragHandlers } from './shared/useEintragHandlers';
import { EintragTable } from './shared/EintragTable';

export const Schritt5_WirtschaftlicherGeschaeftsbetrieb = () => {
  const {
    steuererklaerung,
    errors,
    handleEinnahmeChange,
    handleAusgabeChange,
    addEinnahme,
    addAusgabe,
    removeEinnahme,
    removeAusgabe,
  } = useEintragHandlers('wirtschaftlicherGeschaeftsbetrieb');

  return (
    <div className="space-y-8">
      <EintragTable
        title="Einnahmen im wirtschaftlichen Geschäftsbetrieb"
        tooltip="Alle Einnahmen, die nicht den anderen drei Sphären zugeordnet werden können (z.B. Verkauf von Speisen/Getränken, Werbung, Sponsoring)."
        eintraege={steuererklaerung.wirtschaftlicherGeschaeftsbetrieb.einnahmen}
        errors={errors.einnahmen}
        onEintragChange={handleEinnahmeChange}
        onAddEintrag={addEinnahme}
        onRemoveEintrag={removeEinnahme}
        addButtonText="Einnahme hinzufügen"
      />

      <EintragTable
        title="Ausgaben im wirtschaftlichen Geschäftsbetrieb"
        tooltip="Alle Kosten, die direkt mit den Einnahmen des WGB zusammenhängen (z.B. Wareneinkauf, Personalkosten für den Verkauf)."
        eintraege={steuererklaerung.wirtschaftlicherGeschaeftsbetrieb.ausgaben}
        errors={errors.ausgaben}
        onEintragChange={handleAusgabeChange}
        onAddEintrag={addAusgabe}
        onRemoveEintrag={removeAusgabe}
        addButtonText="Ausgabe hinzufügen"
      />
    </div>
  );
};