import { useState } from "react";
import { ChevronDown, ChevronRight } from "lucide-react";
import CodeBlock from "./CodeBlock";

interface CommandCardProps {
  name: string;
  alias?: string;
  description: string;
  usage?: string;
  flags?: { flag: string; description: string }[];
  examples?: { command: string; description?: string }[];
}

const CommandCard = ({ name, alias, description, usage, flags, examples }: CommandCardProps) => {
  const [open, setOpen] = useState(false);

  return (
    <div className="border border-border rounded-lg overflow-hidden transition-colors hover:border-primary/40">
      <button
        onClick={() => setOpen(!open)}
        className="w-full flex items-center gap-3 px-4 py-3 text-left hover:bg-muted/50 transition-colors"
      >
        {open ? (
          <ChevronDown className="h-4 w-4 text-primary shrink-0" />
        ) : (
          <ChevronRight className="h-4 w-4 text-muted-foreground shrink-0" />
        )}
        <div className="flex items-center gap-2 flex-1 min-w-0">
          <code className="font-mono font-semibold text-sm text-foreground">{name}</code>
          {alias && (
            <span className="text-xs font-mono text-primary bg-primary/10 px-1.5 py-0.5 rounded">{alias}</span>
          )}
        </div>
        <span className="text-sm text-muted-foreground truncate">{description}</span>
      </button>

      {open && (
        <div className="px-4 pb-4 border-t border-border pt-3 space-y-3">
          {usage && <CodeBlock code={usage} />}

          {flags && flags.length > 0 && (
            <div>
              <h4 className="text-xs font-mono font-semibold text-muted-foreground uppercase tracking-wider mb-2">Flags</h4>
              <div className="space-y-1">
                {flags.map((f) => (
                  <div key={f.flag} className="flex gap-4 text-sm">
                    <code className="font-mono text-primary whitespace-nowrap">{f.flag}</code>
                    <span className="text-muted-foreground">{f.description}</span>
                  </div>
                ))}
              </div>
            </div>
          )}

          {examples && examples.length > 0 && (
            <div>
              <h4 className="text-xs font-mono font-semibold text-muted-foreground uppercase tracking-wider mb-2">Examples</h4>
              {examples.map((ex, i) => (
                <div key={i}>
                  {ex.description && <p className="text-sm text-muted-foreground mb-1">{ex.description}</p>}
                  <CodeBlock code={ex.command} />
                </div>
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default CommandCard;
