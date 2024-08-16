using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Baremetal = UnMango.Baremetal;

return await Deployment.RunAsync(() => 
{
    var tee = new Baremetal.Coreutils.Tee("tee", new()
    {
        Args = new Baremetal.Coreutils.Inputs.TeeArgsArgs
        {
            Stdin = "whoops",
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

