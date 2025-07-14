import { useEintragHandlers } from './shared/useEintragHandlers';
import { EintragTable } from './shared/EintragTable';

export const Schritt4_Zweckbetrieb = () => {
  const {
    steuererklaerung,
    errors,
    handleEinnahmeChange,
    handleAusgabeChange,
    addEinnahme,
    addAusgabe,
    removeEinnahme,
    removeAusgabe,
  } = useEintragHandlers('zweckbetrieb');

  return (
    <div className="space-y-8">
      <EintragTable
        title="Einnahmen im Zweckbetrieb"
        tooltip="Einnahmen, die zur Verwirklichung der steuerbegünstigten satzungsmäßigen Zwecke dienen (z.B. Eintrittsgelder für Sportveranstaltungen, Teilnahmegebühren für Kurse)."
        eintraege={steuererklaerung.zweckbetrieb.einnahmen}
        errors={errors.einnahmen}
        onEintragChange={handleEinnahmeChange}
        onAddEintrag={addEinnahme}
        onRemoveEintrag={removeEinnahme}
        addButtonText="Einnahme hinzufügen"
      />

      <EintragTable
        title="Ausgaben im Zweckbetrieb"
        tooltip="Kosten, die direkt mit den Zweckbetrieb-Einnahmen zusammenhängen (z.B. Miete für Sportstätten, Honorare für Übungsleiter)."
        eintraege={steuererklaerung.zweckbetrieb.ausgaben}
        errors={errors.ausgaben}
        onEintragChange={handleAusgabeChange}
        onAddEintrag={addAusgabe}
        onRemoveEintrag={removeAusgabe}
        addButtonText="Ausgabe hinzufügen"
      />
    </div>
  );
};