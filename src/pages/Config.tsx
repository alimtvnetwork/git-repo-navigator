import DocsLayout from "@/components/docs/DocsLayout";
import CodeBlock from "@/components/docs/CodeBlock";

const ConfigPage = () => {
  return (
    <DocsLayout>
      <h1 className="text-3xl font-mono font-bold mb-2">Configuration</h1>
      <p className="text-muted-foreground mb-8">
        Customize gitmap behavior through JSON config files and profiles.
      </p>

      <section className="space-y-8">
        <div>
          <h2 className="text-xl font-mono font-semibold mb-3 text-foreground">config.json</h2>
          <p className="text-muted-foreground mb-3">
            The main config file controls scan defaults. Located at <code className="font-mono text-primary">./data/config.json</code>:
          </p>
          <CodeBlock
            title="data/config.json"
            code={`{
  "defaultMode": "https",
  "defaultOutput": "terminal",
  "outputDir": "gitmap-output",
  "excludeDirs": ["node_modules", ".git", "vendor"],
  "notes": ""
}`}
          />
          <div className="mt-4 space-y-2">
            <div className="flex gap-3 text-sm">
              <code className="font-mono text-primary whitespace-nowrap">defaultMode</code>
              <span className="text-muted-foreground">Clone URL style: <code className="font-mono">https</code> or <code className="font-mono">ssh</code></span>
            </div>
            <div className="flex gap-3 text-sm">
              <code className="font-mono text-primary whitespace-nowrap">defaultOutput</code>
              <span className="text-muted-foreground">Output format: <code className="font-mono">terminal</code>, <code className="font-mono">csv</code>, or <code className="font-mono">json</code></span>
            </div>
            <div className="flex gap-3 text-sm">
              <code className="font-mono text-primary whitespace-nowrap">outputDir</code>
              <span className="text-muted-foreground">Directory for all generated output files</span>
            </div>
            <div className="flex gap-3 text-sm">
              <code className="font-mono text-primary whitespace-nowrap">excludeDirs</code>
              <span className="text-muted-foreground">Directories to skip during recursive scan</span>
            </div>
          </div>
        </div>

        <div>
          <h2 className="text-xl font-mono font-semibold mb-3 text-foreground">git-setup.json</h2>
          <p className="text-muted-foreground mb-3">
            Configure global Git settings applied by <code className="font-mono text-primary">gitmap setup</code>:
          </p>
          <CodeBlock
            title="data/git-setup.json"
            code={`{
  "settings": [
    { "key": "core.autocrlf", "value": "true" },
    { "key": "diff.tool", "value": "vscode" },
    { "key": "merge.tool", "value": "vscode" }
  ]
}`}
          />
        </div>

        <div>
          <h2 className="text-xl font-mono font-semibold mb-3 text-foreground">Profiles</h2>
          <p className="text-muted-foreground mb-3">
            Maintain separate database environments (work, personal, client) using profiles:
          </p>
          <CodeBlock
            code={`# Create a new profile\ngitmap profile create work\n\n# Switch to it\ngitmap profile switch work\n\n# List all profiles\ngitmap profile list\n\n# Compare repos across profiles\ngitmap diff-profiles default work`}
            title="Terminal"
          />
          <p className="text-sm text-muted-foreground mt-2">
            Each profile has its own SQLite database file. The <code className="font-mono text-primary">default</code> profile
            always exists and cannot be deleted. Profile config is stored in{" "}
            <code className="font-mono text-primary">gitmap-output/data/profiles.json</code>.
          </p>
        </div>

        <div>
          <h2 className="text-xl font-mono font-semibold mb-3 text-foreground">CD Defaults</h2>
          <p className="text-muted-foreground mb-3">
            Set default navigation paths for repos cloned to multiple locations:
          </p>
          <CodeBlock
            code={`gitmap cd set-default myrepo C:\\repos\\github\\myrepo\ngitmap cd clear-default myrepo`}
            title="Terminal"
          />
        </div>
      </section>
    </DocsLayout>
  );
};

export default ConfigPage;
