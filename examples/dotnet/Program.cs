using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Baremetal = UnMango.Baremetal;

return await Deployment.RunAsync(() => 
{
    var tee = new Baremetal.Cmd.Tee("tee", new()
    {
        Stdin = "whoops",
        Create = new Baremetal.Cmd.Inputs.TeeOptsArgs
        {
            Files = new[]
            {
                "/tmp/tee/test.txt",
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["stdout"] = tee.Stdout,
    };
});

