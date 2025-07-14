import { Button, Input, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow, Tooltip } from '@heroui/react';
import { Info, Trash2 } from 'lucide-react';
import { type Eintrag } from '@/validations/steuerdaten';
import { cn } from '@/lib/utils';
import { type ErrorState } from './types';

interface EintragTableProps {
  title: string;
  tooltip: string;
  eintraege: Eintrag[];
  errors: { [key: number]: string } | undefined;
  onEintragChange: (index: number, field: keyof Eintrag, value: string | number) => void;
  onAddEintrag: () => void;
  onRemoveEintrag: (index: number) => void;
  addButtonText: string;
}

export const EintragTable = ({
  title,
  tooltip,
  eintraege,
  errors,
  onEintragChange,
  onAddEintrag,
  onRemoveEintrag,
  addButtonText,
}: EintragTableProps) => {
  return (
    <div>
      <div className="flex items-center mb-4">
        <h3 className="text-lg font-medium">{title}</h3>
        <Tooltip content={tooltip}>
          <Info className="h-4 w-4 text-muted-foreground cursor-pointer ml-2" />
        </Tooltip>
      </div>
      <Table>
        <TableHeader>
          <TableColumn>Bezeichnung</TableColumn>
          <TableColumn className="w-[150px]">Betrag (€)</TableColumn>
          <TableColumn className="w-[50px]"> </TableColumn>
        </TableHeader>
        <TableBody>
          {eintraege.map((eintrag, index) => (
            <TableRow key={index}>
              <TableCell>
                <Input
                  type="text"
                  value={eintrag.bezeichnung}
                  onChange={(e) => onEintragChange(index, 'bezeichnung', e.target.value)}
                />
              </TableCell>
              <TableCell>
                <Input
                  type="text"
                  value={(eintrag as any).rawBetrag ?? eintrag.betrag}
                  onChange={(e) => onEintragChange(index, 'betrag', e.target.value)}
                  onBlur={() => {
                    const numValue = parseFloat(
                      ((eintrag as any).rawBetrag ?? eintrag.betrag).toString().replace(',', '.')
                    );
                    if (!isNaN(numValue) && numValue >= 0) {
                      onEintragChange(index, 'betrag', numValue);
                    }
                  }}
                  className={cn(errors?.[index] ? "border-red-500" : "")}
                />
                {errors?.[index] && (
                  <p className="text-xs text-red-600 mt-1">{errors[index]}</p>
                )}
              </TableCell>
              <TableCell>
                <Button variant="ghost" onClick={() => onRemoveEintrag(index)}>
                  <Trash2 className="h-4 w-4" />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Button onClick={onAddEintrag} className="mt-4">
        {addButtonText}
      </Button>
    </div>
  );
};