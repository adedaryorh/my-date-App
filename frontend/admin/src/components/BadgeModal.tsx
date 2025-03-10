import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { bagdes } from "@/lib/constants";
import { Badge } from "@/types/interfaces";
import { Dispatch, SetStateAction } from "react";

export function BadgeModal({ isOpen, onClose, isNotificationModalOpen, setIsNotificationModalOpen, setResponse }: { isOpen: boolean; onClose: () => void; isNotificationModalOpen: boolean; setIsNotificationModalOpen: Dispatch<SetStateAction<boolean>>; setResponse: Dispatch<SetStateAction<string>> }) {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="w-auto">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>Badge selection</DialogTitle>
          </VisuallyHidden>
        </DialogHeader>
        <div className="flex flex-col gap-4 mb-7">
          <Label className="flex flex-col gap-4">
            Select Verification Badge
            <Select>
              <SelectTrigger className="w-[280px]">
                <SelectValue placeholder="Choose a verification badge" />
              </SelectTrigger>
              <SelectContent>
                {bagdes.map((badge: Badge) => (
                  <SelectItem key={badge.id} value={badge.name}>
                    <span className="flex items-center justify-start gap-5">
                      {badge.image} {badge.name}
                    </span>
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </Label>
        </div>
        <DialogFooter>
          <div className="w-full grid grid-cols-1 gap-[18px]">
            <Button
              type="button"
              className="uppercase bg-celebut-gold text-white w-full"
              onClick={() => {
                setResponse("Badge Granted");
                setIsNotificationModalOpen(!isNotificationModalOpen);
              }}
            >
              Grant verification badge
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
