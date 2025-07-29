/**--------------------------------------------------------
 * Author: jiang5630@outlook.com 2025-07-29
 * Description:
 * This file is part of the NTPack project.
 * --------------------------------------------------------
 * 作者：jiang5630@outlook.com  2025年07月29日
 * 描述：通用函数
 --------------------------------------------------------*/

package NTPack

import (
	"os"
)

// GetPid 获取进程ID
func GetPid() int {
	pid := os.Getpid()
	return pid
}
