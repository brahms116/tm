import { createContext, useContext, useEffect, useState } from "react";
import { isAuthenticated as isAuthenticatedQuery } from "@/api";

interface AuthContextType {
  isAuthenticated?: boolean;
  login: (token: string) => Promise<boolean>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

let didAuthInit = false;

export const AuthContextProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isAuthenticated, setAuthenticated] = useState<boolean | undefined>(
    undefined
  );

  useEffect(() => {
    const checkAuth = async () => {
      setAuthenticated(await isAuthenticatedQuery());
    };
    if (didAuthInit) {
      return;
    }
    didAuthInit = true;
    void checkAuth();
  }, []);

  const login = async (token: string): Promise<boolean> => {
    localStorage.setItem("api-key", token);
    const r = await isAuthenticatedQuery();
    setAuthenticated(r);
    return r;
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuthContext = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
