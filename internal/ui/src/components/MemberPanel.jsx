import { useState, useEffect } from "react";
import { Paper, Typography } from "@mui/material";
import MemberForm from "./MemberForm";
import MemberTable from "./MemberTable";
import { AddMember, ListMembers } from "../wailsjs/go/service/DataService";

export default function MemberPanel() {
  const [members, setMembers] = useState([]);

  const fetchMembers = async () => {
    const list = await ListMembers();
    setMembers(list || []);
  };

  useEffect(() => {
    fetchMembers();
  }, []);

  const submit = async (name, email, date, setError) => {
    try {
      await AddMember(name, email, date);
      setError("");
      fetchMembers();
    } catch (err) {
      setError(err.message || "Fehler beim Hinzufügen");
    }
  };

  return (
    <>
      <Paper sx={{ p: 3, mb: 4 }}>
        <Typography variant="h6" component="h2" gutterBottom>
          Neues Mitglied
        </Typography>
        <MemberForm onSubmit={submit} />
      </Paper>
      <Paper>
        <MemberTable members={members} />
      </Paper>
    </>
  );
}
