import Link from "next/link";
import { Menu } from "@/types/interfaces";

export function MenuItem({ item, path, onClick }: { item: Menu; path: string; onClick?: () => void }) {
  return (
    <li className="px-10">
      <Link href={item.url} className={`${path === item.url ? "text-celebut-gold bg-white" : "text-white bg-transparent"} flex gap-4 items-center justify-normal text-sm font-semibold px-6 py-3 rounded-[8px]`} onClick={onClick}>
        <div className="flex justify-center items-center">{item.icon}</div>
        <span>{item.title}</span>
      </Link>
    </li>
  );
}
