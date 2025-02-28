import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { VisuallyHidden } from "@radix-ui/react-visually-hidden";
import { Select } from "@/components/ui/select";
import { Label } from "@/components/ui/label";
// import { Label } from "@/components/ui/label";

export function BadgeModal({ isOpen, onClose }: { isOpen: boolean; onClose: () => void }) {
  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="w-auto">
        <DialogHeader className="w-full">
          <VisuallyHidden>
            <DialogTitle>Badge selection</DialogTitle>
          </VisuallyHidden>
        </DialogHeader>
        <div className="flex flex-col gap-4 mb-7">
          <Label />
          <Select />
        </div>
        <DialogFooter>
          <div className="w-full grid grid-cols-1 gap-[18px]">
            <Button type="button" className="uppercase bg-celebut-gold text-white w-full">
              Grant verification badge
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
