# LeetCode
for learning and job hopping

git config ssh:  
1.ssh-keygen -t rsa -C "helei.cnhk@gmail.com" -f git_id_rsa  
2.ssh config file:  
Host github.com  
HostName github.com  
PreferredAuthentications publickey  
IdentityFile ~/.ssh/git_id_rsa  
3.ssh-add git_id_rsa  
4.git clone ....