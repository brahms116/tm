import { BrowserRouter, Route, Routes } from "react-router";
import { LoginPage } from "./pages/login";
import { DashboardPage } from "./pages/dashboard";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/dashboard" element={<DashboardPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
