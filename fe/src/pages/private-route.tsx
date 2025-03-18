import { useAuthContext } from "@/contexts";
import { PropsWithChildren } from "react";
import { Navigate } from "react-router";

export const PrivateRoute: React.FC<PropsWithChildren> = ({ children }) => {
  const { isAuthenticated } = useAuthContext();
  switch (isAuthenticated) {
    case undefined:
      return <div>Loading...</div>;
    case false:
      return <Navigate to="/login" />;
    case true:
      return <>{children}</>;
  }
};
