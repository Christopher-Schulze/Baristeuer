import { Table, TableHead, TableBody, TableRow, TableCell, Button } from "@mui/material";
import { useTranslation } from "react-i18next";

export default function MemberTable({ members, onEdit, onDelete }) {
  const { t } = useTranslation();
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>{t('member.table.name')}</TableCell>
          <TableCell>{t('member.table.email')}</TableCell>
          <TableCell>{t('member.table.joinDate')}</TableCell>
          <TableCell>{t('member.table.actions')}</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {members.length > 0 ? (
          members.map((m) => (
            <TableRow key={m.id} hover>
              <TableCell>{m.name}</TableCell>
              <TableCell>{m.email}</TableCell>
              <TableCell>{m.joinDate}</TableCell>
              <TableCell>
                {onEdit && (
                  <Button size="small" onClick={() => onEdit(m)}>
                    {t('edit')}
                  </Button>
                )}
                <Button size="small" color="error" onClick={() => onDelete(m.id)}>
                  {t('delete')}
                </Button>
              </TableCell>
            </TableRow>
          ))
        ) : (
          <TableRow>
            <TableCell colSpan={4} align="center">
              {t('member.table.empty')}
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
