import { createContext, useContext, useEffect, useState } from "react";

interface AuthContextType {
  isAuthenticated?: boolean;
  login: (token: string) => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

let didAuthInit = false;

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isAuthenticated, setAuthenticated] = useState<boolean | undefined>(
    undefined
  );

  useEffect(() => {
    if (didAuthInit) {
      return;
    }
    didAuthInit = true;
    const storedApiKey = localStorage.getItem("api-key");
    if (storedApiKey) {
      setAuthenticated(true);
    }
  }, []);

  const login = async (token: string) => {
    localStorage.setItem("api-key", token);
    setAuthenticated(true);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, login }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
