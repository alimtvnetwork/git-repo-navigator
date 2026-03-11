import { format } from "date-fns";
import {
  FolderGit2,
  Code2,
  Braces,
  Cpu,
  Hash,
  MapPin,
  FileText,
  FileCode,
  Clock,
  FolderOpen,
} from "lucide-react";
import { Badge } from "@/components/ui/badge";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Separator } from "@/components/ui/separator";
import type { DetectedProject, ProjectType } from "@/components/projects/types";

const PROJECT_TYPES: Record<ProjectType, { label: string; color: string; icon: typeof Code2 }> = {
  go: { label: "Go", color: "bg-cyan-500/15 text-cyan-700 dark:text-cyan-400 border-cyan-500/30", icon: Code2 },
  node: { label: "Node.js", color: "bg-emerald-500/15 text-emerald-700 dark:text-emerald-400 border-emerald-500/30", icon: Braces },
  react: { label: "React", color: "bg-sky-500/15 text-sky-700 dark:text-sky-400 border-sky-500/30", icon: Braces },
  cpp: { label: "C++", color: "bg-violet-500/15 text-violet-700 dark:text-violet-400 border-violet-500/30", icon: Cpu },
  csharp: { label: "C#", color: "bg-purple-500/15 text-purple-700 dark:text-purple-400 border-purple-500/30", icon: Hash },
};

const DetailRow = ({ icon: Icon, label, value }: { icon: typeof MapPin; label: string; value: string }) => (
  <div className="flex items-start gap-3 py-2">
    <Icon className="h-4 w-4 text-muted-foreground mt-0.5 shrink-0" />
    <div className="min-w-0">
      <span className="text-xs text-muted-foreground block">{label}</span>
      <span className="font-mono text-sm text-foreground break-all">{value}</span>
    </div>
  </div>
);

interface Props {
  project: DetectedProject | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

const ProjectDetailDialog = ({ project, open, onOpenChange }: Props) => {
  if (!project) return null;

  const typeConfig = PROJECT_TYPES[project.projectType];
  const TypeIcon = typeConfig.icon;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-lg max-h-[85vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle className="flex items-center gap-2 font-mono text-lg">
            <Badge variant="outline" className={`${typeConfig.color} font-mono text-xs gap-1 border`}>
              <TypeIcon className="h-3 w-3" />
              {typeConfig.label}
            </Badge>
            <span className="truncate">{project.projectName}</span>
          </DialogTitle>
        </DialogHeader>

        <div className="space-y-1">
          <DetailRow icon={FolderGit2} label="Repository" value={project.repoName} />
          <DetailRow icon={FolderOpen} label="Absolute Path" value={project.absolutePath} />
          <DetailRow icon={MapPin} label="Repo Path" value={project.repoPath} />
          <DetailRow
            icon={MapPin}
            label="Relative Path"
            value={project.relativePath === "." ? "(root)" : project.relativePath}
          />
          <DetailRow icon={FileText} label="Primary Indicator" value={project.primaryIndicator} />
          <DetailRow
            icon={Clock}
            label="Detected At"
            value={format(new Date(project.detectedAt), "PPpp")}
          />
        </div>

        {project.goMetadata && (
          <>
            <Separator />
            <div className="space-y-3">
              <h3 className="font-mono text-sm font-semibold text-foreground flex items-center gap-2">
                <Code2 className="h-4 w-4 text-primary" />
                Go Metadata
              </h3>
              <div className="grid grid-cols-2 gap-3">
                <div className="rounded-md bg-muted p-3">
                  <span className="text-xs text-muted-foreground block">Module</span>
                  <span className="font-mono text-sm text-foreground break-all">{project.goMetadata.moduleName}</span>
                </div>
                <div className="rounded-md bg-muted p-3">
                  <span className="text-xs text-muted-foreground block">Go Version</span>
                  <span className="font-mono text-sm text-foreground">{project.goMetadata.goVersion}</span>
                </div>
              </div>
              {project.goMetadata.runnables.length > 0 && (
                <div>
                  <span className="text-xs text-muted-foreground block mb-2">Runnable Entry Points</span>
                  <div className="space-y-1.5">
                    {project.goMetadata.runnables.map((r) => (
                      <div key={r.name} className="flex items-center gap-2 rounded-md bg-muted px-3 py-2">
                        <FileCode className="h-3.5 w-3.5 text-primary shrink-0" />
                        <span className="font-mono text-sm font-medium text-foreground">{r.name}</span>
                        <span className="font-mono text-xs text-muted-foreground ml-auto truncate">{r.relativePath}</span>
                      </div>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </>
        )}

        {project.csharpMetadata && (
          <>
            <Separator />
            <div className="space-y-3">
              <h3 className="font-mono text-sm font-semibold text-foreground flex items-center gap-2">
                <Hash className="h-4 w-4 text-primary" />
                C# Metadata
              </h3>
              <div className="grid grid-cols-2 gap-3">
                <div className="rounded-md bg-muted p-3">
                  <span className="text-xs text-muted-foreground block">Solution</span>
                  <span className="font-mono text-sm text-foreground">{project.csharpMetadata.slnName}</span>
                </div>
                <div className="rounded-md bg-muted p-3">
                  <span className="text-xs text-muted-foreground block">SDK Version</span>
                  <span className="font-mono text-sm text-foreground">{project.csharpMetadata.sdkVersion}</span>
                </div>
              </div>
              {project.csharpMetadata.projectFiles.length > 0 && (
                <div>
                  <span className="text-xs text-muted-foreground block mb-2">Project Files</span>
                  <div className="space-y-1.5">
                    {project.csharpMetadata.projectFiles.map((f) => (
                      <div key={f.fileName} className="flex items-center gap-2 rounded-md bg-muted px-3 py-2">
                        <FileCode className="h-3.5 w-3.5 text-primary shrink-0" />
                        <span className="font-mono text-sm text-foreground">{f.fileName}</span>
                        <Badge variant="outline" className="text-[10px] px-1.5 py-0 ml-auto">{f.targetFramework}</Badge>
                        <span className="text-xs text-muted-foreground">{f.outputType}</span>
                      </div>
                    ))}
                  </div>
                </div>
              )}
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};

export default ProjectDetailDialog;
