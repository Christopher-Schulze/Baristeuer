import { Table, TableHead, TableBody, TableRow, TableCell, Button } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function ExpenseTable({ expenses, onEdit, onDelete }) {
  const { t } = useTranslation();
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>{t('expense.table.description')}</TableCell>
          <TableCell align="right">{t('expense.table.amount')}</TableCell>
          <TableCell>{t('expense.table.actions')}</TableCell>
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
                  {t('edit')}
                </Button>
                <Button size="small" color="error" onClick={() => onDelete(e.id)}>
                  {t('delete')}
                </Button>
              </TableCell>
            </TableRow>
          ))
        ) : (
          <TableRow>
            <TableCell colSpan={3} align="center">
              {t('expense.table.empty')}
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
