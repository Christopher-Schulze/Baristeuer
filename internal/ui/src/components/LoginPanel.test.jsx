import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import "@testing-library/jest-dom";
import { vi } from "vitest";
import LoginPanel from "./LoginPanel";
import "../i18n";

vi.mock(
  "../wailsjs/go/service/DataService",
  () => ({
    Login: vi.fn().mockResolvedValue(),
    Register: vi.fn().mockResolvedValue(),
  }),
  { virtual: true },
);

import { Login, Register } from "../wailsjs/go/service/DataService";

test("calls login function", async () => {
  render(<LoginPanel onLoggedIn={() => {}} />);
  fireEvent.change(screen.getByLabelText(/Benutzer/i), {
    target: { value: "a" },
  });
  fireEvent.change(screen.getByLabelText(/Passwort/i), {
    target: { value: "b" },
  });
  fireEvent.click(screen.getByRole("button", { name: /Anmelden/i }));
  await waitFor(() => expect(Login).toHaveBeenCalledWith("a", "b"));
});
