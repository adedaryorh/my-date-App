"use client";

import { InputOTP, InputOTPGroup, InputOTPSlot } from "@/components/ui/input-otp";
import { REGEXP_ONLY_DIGITS } from "input-otp";
import { Button } from "@/components/ui/button";
import { useEffect, useState } from "react";

export default function OtpForm() {
  const [count, setCount] = useState<number>(59);
  useEffect(() => {
    if (count === 0) return;

    const timer = setInterval(() => {
      setCount(count - 1);
    }, 1000);

    return () => {
      clearTimeout(timer);
    };
  }, [count]);

  const handleLogin = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
  };
  return (
    <div className="flex flex-col items-center justify-center w-3/4">
      <form className="flex flex-col items-center w-full gap-12">
        <div className="flex flex-col items-center w-full gap-6">
          <InputOTP maxLength={6} pattern={REGEXP_ONLY_DIGITS}>
            <InputOTPGroup className="outline-none focus:outline-none">
              <InputOTPSlot index={0} className="h-14 w-16 bg-gray-100" />
              <InputOTPSlot index={1} className="h-14 w-16 bg-gray-100" />
              <InputOTPSlot index={2} className="h-14 w-16 bg-gray-100" />
              <InputOTPSlot index={3} className="h-14 w-16 bg-gray-100" />
              <InputOTPSlot index={4} className="h-14 w-16 bg-gray-100" />
              <InputOTPSlot index={5} className="h-14 w-16 bg-gray-100" />
            </InputOTPGroup>
          </InputOTP>
        </div>
        <Button type="submit" className="w-full">
          Verify
        </Button>
        <p className="flex justify-center items-center gap-2 font-semibold text-base">
          00:{count < 10 ? count.toString().padStart(2, "0") : count}{" "}
          <button disabled={count > 0} type="button" className={`${count === 0 ? "text-black" : "text-[#C4C4C4]"} bg-none font-normal`} onClick={handleLogin}>
            Resend code
          </button>
        </p>
      </form>
    </div>
  );
}
