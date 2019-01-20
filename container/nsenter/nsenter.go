package nsenter

/*
#define _GNU_SOURCE
#include <unistd.h>
#include <errno.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>

__attribute__((constructor)) void enter_namespace(void) {
	char *dockerc_pid;
	dockerc_pid = getenv("dockerc_pid");
	if (dockerc_pid) {
		//fprintf(stdout, "got dockerc_pid=%s\n", dockerc_pid);
	} else {
		//fprintf(stdout, "missing dockerc_pid env skip nsenter");
		return;
	}
	char *dockerc_cmd;
	dockerc_cmd = getenv("dockerc_cmd");
	if (dockerc_cmd) {
		//fprintf(stdout, "got dockerc_cmd=%s\n", dockerc_cmd);
	} else {
		//fprintf(stdout, "missing dockerc_cmd env skip nsenter");
		return;
	}
	int i;
	char nspath[1024];
	char *namespaces[] = { "ipc", "uts", "net", "pid", "mnt" };

	for (i=0; i<5; i++) {
		sprintf(nspath, "/proc/%s/ns/%s", dockerc_pid, namespaces[i]);
		int fd = open(nspath, O_RDONLY);

		if (setns(fd, 0) == -1) {
			//fprintf(stderr, "setns on %s namespace failed: %s\n", namespaces[i], strerror(errno));
		} else {
			//fprintf(stdout, "setns on %s namespace succeeded\n", namespaces[i]);
		}
		close(fd);
	}
	int res = system(dockerc_cmd);
	exit(0);
	return;
}
*/
import "C"
