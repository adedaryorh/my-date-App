import { MenuItems } from "@/lib/constants";
import { MenuItem } from "@/components/MenuItem";
import Link from "next/link";

export default function MenuList({ path, handleNavigation }: { path: string; handleNavigation?: () => void }) {
  return (
    <ul className="py-8">
      {MenuItems.map((item) => (
        <MenuItem key={item.url} item={item} path={path} onClick={handleNavigation} />
      ))}
      <li className="border-t px-10 mt-4 pt-4">
        <Link href="/user-configuration" className={`${path === "/user-configuration" ? "text-celebut-gold bg-white" : "text-white bg-transparent"} flex gap-4 items-center justify-normal text-sm font-semibold px-6 py-3 rounded-[8px]`}>
          <div className="flex justify-center items-center">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12 23C6.443 21.765 2 16.522 2 11V5L12 1L22 5V11C22 16.524 17.557 21.765 12 23ZM4 6V11C4.05715 13.3121 4.87036 15.5418 6.31518 17.3479C7.75999 19.1539 9.75681 20.4367 12 21C14.2432 20.4367 16.24 19.1539 17.6848 17.3479C19.1296 15.5418 19.9429 13.3121 20 11V6L12 3L4 6Z" fill="white" />
              <path d="M12 11C13.3807 11 14.5 9.88071 14.5 8.5C14.5 7.11929 13.3807 6 12 6C10.6193 6 9.5 7.11929 9.5 8.5C9.5 9.88071 10.6193 11 12 11Z" fill="white" />
              <path d="M7 15C7.49273 15.8983 8.21539 16.6496 9.09398 17.1767C9.97256 17.7039 10.9755 17.988 12 18C13.0245 17.988 14.0274 17.7039 14.906 17.1767C15.7846 16.6496 16.5073 15.8983 17 15C16.975 13.104 13.658 12 12 12C10.333 12 7.025 13.104 7 15Z" fill="white" />
            </svg>
          </div>
          <span>Super Admin</span>
        </Link>
      </li>
      <li className="border-t px-10 mt-4 py-4">
        <button type="button" className="text-white bg-transparent flex gap-4 items-center justify-normal text-sm font-semibold px-6 py-3 rounded-[8px]">
          <div className="flex justify-center items-center">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M12.0003 2V12M18.4003 6.6C19.6569 7.85711 20.5132 9.45817 20.8611 11.2013C21.209 12.9444 21.0329 14.7515 20.3551 16.3947C19.6774 18.0379 18.5282 19.4436 17.0525 20.4345C15.5769 21.4254 13.8408 21.9572 12.0634 21.9627C10.2859 21.9683 8.54654 21.4474 7.06471 20.4658C5.58288 19.4841 4.42491 18.0856 3.73684 16.4467C3.04876 14.8078 2.8614 13.0019 3.19837 11.2566C3.53533 9.51135 4.38155 7.90495 5.63029 6.64" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
            </svg>
          </div>
          <span>Logout</span>
        </button>
      </li>
    </ul>
  );
}
