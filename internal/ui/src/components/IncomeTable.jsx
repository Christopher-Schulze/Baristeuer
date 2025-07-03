import { Table, TableHead, TableBody, TableRow, TableCell, Button } from "@mui/material";

export default function IncomeTable({ incomes, onEdit, onDelete }) {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Quelle</TableCell>
          <TableCell align="right">Betrag (€)</TableCell>
          <TableCell>Aktionen</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {incomes.length > 0 ? (
          incomes.map((i) => (
            <TableRow key={i.id} hover>
              <TableCell>{i.source}</TableCell>
              <TableCell align="right">{i.amount.toFixed(2)}</TableCell>
              <TableCell>
                <Button size="small" onClick={() => onEdit(i)}>
                  Bearbeiten
                </Button>
                <Button size="small" color="error" onClick={() => onDelete(i.id)}>
                  Löschen
                </Button>
              </TableCell>
            </TableRow>
          ))
        ) : (
          <TableRow>
            <TableCell colSpan={3} align="center">
              Keine Einnahmen vorhanden
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
