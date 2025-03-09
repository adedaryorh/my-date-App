import { Dialog, DialogClose, DialogContent, DialogTitle } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { X } from "lucide-react";
import { useState } from "react";
import { FilterModalProps } from "@/types/interfaces";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";

export default function FilterModal({ open, onOpenChange, onApplyFilters }: FilterModalProps) {
  const [selectedStatus, setSelectedStatus] = useState<string[]>([]);

  const [continent, setContinent] = useState<string>();
  const [country, setCountry] = useState<string>();
  const [city, setCity] = useState<string>();
  const [town, setTown] = useState<string>();

  const statusOptions = ["Active", "Warned", "Suspended", "Removed", "Restored", "Inactive"];

  // data for dropdowns
  const continents = ["Africa", "Asia", "Europe", "North America", "South America", "Australia", "Antarctica"];
  const countries = ["United States", "Canada", "United Kingdom", "Germany", "France", "Japan", "Australia"];
  const cities = ["New York", "London", "Paris", "Tokyo", "Sydney", "Berlin", "Toronto"];
  const towns = ["Brooklyn", "Westminster", "Montmartre", "Shibuya", "Bondi", "Kreuzberg", "Yorkville"];

  const toggleStatus = (status: string) => {
    if (selectedStatus.includes(status)) {
      setSelectedStatus(selectedStatus.filter((s) => s !== status));
    } else {
      setSelectedStatus([...selectedStatus, status]);
    }
  };

  const handleApply = () => {
    const filters = { status: selectedStatus };
    if (onApplyFilters) {
      onApplyFilters(filters);
    }
    onOpenChange(false);
    console.log("Applied filters:", { status: selectedStatus });
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-[500px] md:max-w-xl p-0 border-none [&>button]:hidden" onInteractOutside={(e) => e.preventDefault()}>
        <div className="p-6">
          <div className="flex justify-between items-center mb-6">
            <div className="flex items-center">
              <DialogClose className="p-0 h-auto bg-transparent hover:bg-transparent">
                <X className="w-6 h-6 text-[#000000] mr-4" />
              </DialogClose>
              <DialogTitle className="text-3xl font-bold text-[#000000] m-0 p-0">Filter</DialogTitle>
            </div>
            <Button className="bg-[#f5bd4b] hover:bg-[#f5bd4b]/90 text-[#ffffff] px-8 py-2 rounded-lg font-bold" onClick={handleApply}>
              APPLY
            </Button>
          </div>

          <div className="border-t border-[#e6e6e6] py-6">
            <h2 className="text-base font-semibold text-[#000000] mb-4">Status</h2>
            <div className="flex flex-wrap gap-3">
              {statusOptions.map((status) => (
                <button type="button" key={status} onClick={() => toggleStatus(status)} className={cn("text-sm font-normal p-2 rounded-lg border border-[#000000] text-[#000000]", selectedStatus.includes(status) && "bg-[#e6e6e6]")}>
                  {status}
                </button>
              ))}
            </div>
          </div>

          <div className="border-t border-[#e6e6e6] py-6">
            <h2 className="text-base font-semibold text-[#000000] mb-4">Location</h2>
            <div className="flex flex-col gap-4">
              <Select value={continent} onValueChange={setContinent}>
                <SelectTrigger className="border-[#000000] rounded-lg h-12 px-4 py-3">
                  <SelectValue placeholder="Continent" />
                </SelectTrigger>
                <SelectContent>
                  {continents.map((item) => (
                    <SelectItem key={item} value={item}>
                      {item}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select value={country} onValueChange={setCountry}>
                <SelectTrigger className="border-[#000000] rounded-lg h-12 px-4 py-3">
                  <SelectValue placeholder="Country" />
                </SelectTrigger>
                <SelectContent>
                  {countries.map((item) => (
                    <SelectItem key={item} value={item}>
                      {item}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select value={city} onValueChange={setCity}>
                <SelectTrigger className="border-[#000000] rounded-lg h-12 px-4 py-3">
                  <SelectValue placeholder="City" />
                </SelectTrigger>
                <SelectContent>
                  {cities.map((item) => (
                    <SelectItem key={item} value={item}>
                      {item}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select value={town} onValueChange={setTown}>
                <SelectTrigger className="border-[#000000] rounded-lg h-12 px-4 py-3">
                  <SelectValue placeholder="Town" />
                </SelectTrigger>
                <SelectContent>
                  {towns.map((item) => (
                    <SelectItem key={item} value={item}>
                      {item}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
