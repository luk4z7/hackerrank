class Node
	attr_reader :value

	attr_accessor :right, :left

	 def initialize(value)
	  @value = value
	  @left =nil
	  @right =nil
	 end
end

def choice(n)
	if n == 1
		return 2,3
	elsif n == 2
		return 1,3
	else
		return 1,2
	end		
end

def pushNode(node,first,second)
	if node.left
		pushNode(node.left,first,second)
	else
		node.left = Node.new(second)
	end

	if node.right
		pushNode(node.right, first,second)
	else
		node.right = Node.new(first)
	end

end

def traverse(node)
	if node == nil
		return
	end

	if node.left
		traverse(node.left)
	end
  
	if node.right
		traverse(node.right)
	end
end

def minimumMovement(obstacleLanes)
    # Write your code here
    arr = Array.new
    obstacleLanes.each do |lane|
        arr << choice(lane)	
    end
    root = Node.new(2);
	  arr.each do |inner| 
		pushNode(root, inner[0], inner[1])
	end
	traverse(root)
end

minimumMovement([2,1,2])

# minimumMovement([2,1,3,2])
