"use client";

import AuthForm from "@/components/AuthForm";
import LoginForm from "@/components/LoginForm";
import Image from "next/image";

export default function Login() {
  return (
    <AuthForm>
      <>
        <div className="flex flex-col items-center">
          <Image src="/img/logo.png" alt="Celebut" width={250} height={250} className="w-auto" />
        </div>
        <LoginForm />
      </>
    </AuthForm>
  );
}
