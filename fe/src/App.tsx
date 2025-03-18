import { BrowserRouter, Route, Routes } from "react-router";
import { LoginPage } from "@/pages/login";
import { DashboardPage } from "@/pages/dashboard";
import { AuthContextProvider } from "@/contexts";
import { Toaster } from "./components/ui/sonner";
import { PrivateRoute } from "./pages/private-route";

function App() {
  return (
    <AuthContextProvider>
      <Toaster />
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route
            path="/dashboard"
            element={
              <PrivateRoute>
                <DashboardPage />
              </PrivateRoute>
            }
          />
        </Routes>
      </BrowserRouter>
    </AuthContextProvider>
  );
}

export default App;
