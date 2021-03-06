// Copyright (c) 2018 SUSE LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

// zypper-docker list-patches [flags] <image>
func listPatchesCmd(ctx *cli.Context) {
	imageID := ctx.Args().First()
	err := listPatches(imageID, ctx)
	exitOnError(imageID, "zypper lp", err)
}

// zypper-docker list-patches-container [flags] <container>
func listPatchesContainerCmd(ctx *cli.Context) {
	imageID, err := commandInContainer(listPatches, ctx)
	exitOnError(imageID, "zypper lp", err)
}

// listParches calls the `zypper lp` command for the given image and the given
// arguments.
func listPatches(image string, ctx *cli.Context) error {
	if image == "" {
		logAndFatalf("Error: no image name specified.\n")
		exitWithCode(1)
	}

	if severity := ctx.String("severity"); severity != "" {
		if ok, err := supportsSeverityFlag(image); !ok {
			if err == nil {
				log.Println("the --severity flag is only available for zypper versions >= 1.12.6")
				fmt.Println("the --severity flag is only available for zypper versions >= 1.12.6")
			} else {
				log.Println(err)
				fmt.Println(err)
			}
			exitWithCode(1)
		}
	}

	err := runStreamedCommand(
		image,
		cmdWithFlags("lp", ctx, []string{}, []string{"base"}), true)
	return err
}

// zypper-docker patch [flags] image
func patchCmd(ctx *cli.Context) {
	updatePatchCmd("patch", ctx)
}
