import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { useAuthContext } from "@/contexts";
import { Loader2 } from "lucide-react";
import { toast } from "sonner";
import { useNavigate } from "react-router";

export function LoginForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"div">) {
  const { isAuthenticated, login } = useAuthContext();
  const [isLoginLoading, setIsLoginLoading] = useState(false);
  const [apiKeyInput, setApiKeyInput] = useState("");
  const navigate = useNavigate();

  // TODO: Fix with spinner
  if (isAuthenticated) {
    navigate("/dashboard");
    return;
  }

  const handleFormSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoginLoading(true);
    const r = await login(apiKeyInput);
    if (!r) {
      toast.error("Invalid api key");
      setIsLoginLoading(false);
    } else {
      navigate("/dashboard");
    }
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
          <CardDescription>Enter your api key.</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleFormSubmit}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Api key</Label>
                </div>
                <Input
                  id="password"
                  type="password"
                  required
                  onChange={(e) => setApiKeyInput(e.target.value)}
                  value={apiKeyInput}
                />
              </div>
              <Button
                type="submit"
                className="w-full"
                disabled={isLoginLoading}
              >
                {isLoginLoading ? <Loader2 className="animate-spin" /> : null}
                Login
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
