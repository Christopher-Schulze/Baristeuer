import { Table, TableHead, TableBody, TableRow, TableCell, Button } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function IncomeTable({ incomes, onEdit, onDelete }) {
  const { t } = useTranslation();
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>{t('income.table.source')}</TableCell>
          <TableCell align="right">{t('income.table.amount')}</TableCell>
          <TableCell>{t('income.table.actions')}</TableCell>
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
                  {t('edit')}
                </Button>
                <Button size="small" color="error" onClick={() => onDelete(i.id)}>
                  {t('delete')}
                </Button>
              </TableCell>
            </TableRow>
          ))
        ) : (
          <TableRow>
            <TableCell colSpan={3} align="center">
              {t('income.table.empty')}
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
