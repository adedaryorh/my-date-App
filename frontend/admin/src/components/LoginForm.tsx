import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function LoginForm({ setStep }: { setStep: (step: string) => void }) {
  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    setStep("otp");
  };
  return (
    <div className="flex flex-col items-center justify-center w-3/4">
      <form className="flex flex-col items-center w-full gap-12">
        <div className="flex flex-col items-center w-full gap-6">
          <Input type="email" placeholder="Email/Phone Number" className="bg-gray-100" />
          <Input type="password" placeholder="Password" className="bg-gray-100" />
        </div>
        <Button type="submit" className="w-full" onClick={handleClick}>
          LOGIN
        </Button>
      </form>
    </div>
  );
}
