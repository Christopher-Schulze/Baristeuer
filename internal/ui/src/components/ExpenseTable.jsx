import { Table, TableHead, TableBody, TableRow, TableCell, Button } from "@mui/material";

export default function ExpenseTable({ expenses, onEdit, onDelete }) {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Beschreibung</TableCell>
          <TableCell align="right">Betrag (€)</TableCell>
          <TableCell>Aktionen</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {expenses.length > 0 ? (
          expenses.map((e) => (
            <TableRow key={e.id} hover>
              <TableCell>{e.description}</TableCell>
              <TableCell align="right">{e.amount.toFixed(2)}</TableCell>
              <TableCell>
                <Button size="small" onClick={() => onEdit(e)}>
                  Bearbeiten
                </Button>
                <Button size="small" color="error" onClick={() => onDelete(e.id)}>
                  Löschen
                </Button>
              </TableCell>
            </TableRow>
          ))
        ) : (
          <TableRow>
            <TableCell colSpan={3} align="center">
              Keine Ausgaben vorhanden
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
