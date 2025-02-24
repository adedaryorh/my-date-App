"use client";

import MenuList from "@/components/MenuList";
import { headerIcons } from "@/lib/constants";
import { Menu, X } from "lucide-react";
import { usePathname } from "next/navigation";
import { motion, AnimatePresence } from "motion/react";
import { useState } from "react";

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
      <aside className="hidden lg:block w-[332px] bg-celebut-gold overflow-scroll">
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
        <header className="flex items-center justify-between px-10 py-9">
          <div className="flex items-center gap-4">
            <button className="lg:hidden" onClick={toggleMenu} type="button" title="Open menu">
              <Menu color="#F5BD4B" size={30} />
            </button>
            <h1 className="capitalize text-4xl font-bold">{heading}</h1>
          </div>
          <div className="flex items-center gap-6">
            {headerIcons.map((headerIcon) => (
              <button key={headerIcon.title} className="p-2">
                {headerIcon.icon}
              </button>
            ))}
          </div>
        </header>
        <main className="p-10">{children}</main>
      </div>
    </section>
  );
}
