"use client";

import AuthForm from "@/components/AuthForm";
import LoginForm from "@/components/LoginForm";
import OtpForm from "@/components/OtpForm";
import Image from "next/image";
import { useState } from "react";

export default function Login() {
  const [step, setStep] = useState("login");
  return (
    <AuthForm>
      <div className="flex flex-col items-center">
        <Image src="/img/logo.png" alt="Celebut" width={250} height={250} className="w-auto" />
      </div>
      {step === "login" ? <LoginForm setStep={setStep} /> : <OtpForm />}
    </AuthForm>
  );
}
