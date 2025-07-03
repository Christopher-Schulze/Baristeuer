import { Table, TableHead, TableRow, TableCell, TableBody } from "@mui/material";

export default function MemberTable({ members }) {
  return (
    <Table size="small">
      <TableHead>
        <TableRow>
          <TableCell>Name</TableCell>
          <TableCell>Email</TableCell>
          <TableCell>Beitritt</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {members.map((m) => (
          <TableRow key={m.id}>
            <TableCell>{m.name}</TableCell>
            <TableCell>{m.email}</TableCell>
            <TableCell>{m.join_date || m.joinDate}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
