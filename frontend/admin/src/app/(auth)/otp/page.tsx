"use client";

import AuthForm from "@/components/AuthForm";
import OtpForm from "@/components/OtpForm";

export default function Login() {
  return (
    <AuthForm>
      <div className="flex flex-col items-center gap-4 pt-16">
        <h1 className="text-4xl font-semibold">Email Verification</h1>
        <p className="font-normal text-base">A verification code has been sent to your email</p>
      </div>
      <OtpForm />
    </AuthForm>
  );
}
