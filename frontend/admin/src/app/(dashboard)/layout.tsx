"use client";

import MenuList from "@/components/MenuList";
import { headerIcons } from "@/lib/constants";
import { Menu, X } from "lucide-react";
import { usePathname } from "next/navigation";
import { motion, AnimatePresence } from "motion/react";
import { useState } from "react";
import Image from "next/image";

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const heading = pathname.replaceAll("-", " ").substring(1);

  const [isOpen, setIsOpen] = useState<boolean>(false);

  const toggleMenu = () => {
    setIsOpen(!isOpen);
    // Prevent scrolling when menu is open
    document.body.style.overflow = !isOpen ? "hidden" : "unset";
  };
  // Close menu on navigation
  const handleNavigation = () => {
    setIsOpen(false);
    document.body.style.overflow = "unset";
  };
  return (
    <section className="flex h-screen relative">
      <aside className="hidden lg:block w-[332px] bg-celebut-gold overflow-scroll scrollbar-hide">
        <div className="flex items-center justify-between px-5 lg:px-10 pt-[30px]">
          <figure className="w-fit h-fit rounded-full border-2 border-white">
            <Image src="/img/user.jpeg" alt="user" width={50} height={50} className="rounded-full w-[50px] h-[50px] object-cover" />
          </figure>
          <div>
            <p className="font-bold text-sm text-white">Oluwaferanmi Oluwatobi</p>
            <p className="font-normal text-xs text-white">Admin</p>
          </div>
        </div>
        <MenuList path={pathname} />
      </aside>

      {/* Mobile nav */}
      <AnimatePresence>
        {isOpen && (
          <motion.aside className="h-screen w-full bg-celebut-gold overflow-scroll absolute" initial={{ opacity: 0, x: -500 }} animate={{ opacity: 1, x: 0 }} exit={{ opacity: 0, x: -500 }} transition={{ duration: 0.5 }}>
            <div className="flex justify-end p-5 fixed top-0 right-0">
              <X size={30} color="#fff" onClick={toggleMenu} />
            </div>
            <div className="mt-10">
              <MenuList handleNavigation={handleNavigation} path={pathname} />
            </div>
          </motion.aside>
        )}
      </AnimatePresence>

      <div className="flex-1 overflow-auto bg-white border">
        <header className="flex items-center justify-between px-5 lg:px-10 py-10 border-b">
          <div className="flex items-center gap-4">
            <button className="lg:hidden" onClick={toggleMenu} type="button" title="Open menu">
              <Menu color="#F5BD4B" size={30} />
            </button>
            <h1 className="capitalize text-xl lg:text-4xl font-bold">{heading}</h1>
          </div>
          <div className="flex items-center gap-1 lg:gap-6">
            {headerIcons.map((headerIcon) => (
              <button type="button" key={headerIcon.title} className="p-2">
                {headerIcon.icon}
              </button>
            ))}
          </div>
        </header>
        <main className="p-5 lg:p-10">{children}</main>
      </div>
    </section>
  );
}
