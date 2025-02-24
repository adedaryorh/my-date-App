type Props = {
  children: React.ReactElement;
};

export default function AuthForm({ children }: Props) {
  return (
    <section className="grid grid-cols-1 lg:grid-cols-9 items-center h-screen">
      <div className="relative hidden lg:block md:col-span-5 bg-white h-full" style={{ backgroundImage: "url('/img/login-background.jpg')", backgroundSize: "cover", backgroundPosition: "center" }}>
        <div className="absolute inset-0 bg-celebut-overlay opacity-50 h-full w-full top-0 left-0"></div>
        <p className="absolute text-white font-extrabold text-5xl leading-[67.2px] text-center z-10 bottom-[120px] left-1/2 -translate-x-1/2">Celebrate Your Loved Ones</p>
      </div>
      <div className="lg:col-span-4 bg-white h-full flex flex-col items-center gap-16 justify-start pt-16">{children}</div>
    </section>
  );
}
