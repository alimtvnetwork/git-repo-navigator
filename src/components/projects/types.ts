export type ProjectType = "go" | "node" | "react" | "cpp" | "csharp";

export interface DetectedProject {
  id: string;
  repoName: string;
  projectType: ProjectType;
  projectName: string;
  absolutePath: string;
  repoPath: string;
  relativePath: string;
  primaryIndicator: string;
  detectedAt: string;
  goMetadata?: {
    moduleName: string;
    goVersion: string;
    runnables: { name: string; relativePath: string }[];
  };
  csharpMetadata?: {
    slnName: string;
    sdkVersion: string;
    projectFiles: { fileName: string; targetFramework: string; outputType: string }[];
  };
}
