import { useState, useCallback } from "react";
import { Copy, Check, Terminal } from "lucide-react";

interface InstallBlockProps {
  command: string;
}

const InstallBlock = ({ command }: InstallBlockProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = useCallback(() => {
    navigator.clipboard.writeText(command);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  }, [command]);

  return (
    <div
      onClick={handleCopy}
      className="inline-flex items-center gap-3 px-5 py-3 rounded-lg bg-terminal border border-border cursor-pointer hover:border-primary/40 transition-colors group"
    >
      <Terminal className="h-4 w-4 text-primary" />
      <code className="font-mono text-sm text-terminal-foreground">{command}</code>
      <span className="text-muted-foreground group-hover:text-foreground transition-colors">
        {copied ? <Check className="h-4 w-4 text-primary" /> : <Copy className="h-4 w-4" />}
      </span>
    </div>
  );
};

export default InstallBlock;
