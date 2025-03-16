import { Flex, Icon, Text } from '@chakra-ui/react';
import { FaShoppingBag } from 'react-icons/fa';
import { Link as RouterLink } from 'react-router-dom';

const Logo = ({ color = 'brand.500', size = 'md' }) => {
  const logoSizes = {
    sm: {
      fontSize: 'xl',
      iconSize: 5,
    },
    md: {
      fontSize: '2xl',
      iconSize: 6,
    },
    lg: {
      fontSize: '3xl',
      iconSize: 8,
    },
  };

  const { fontSize, iconSize } = logoSizes[size];

  return (
    <Flex
      as={RouterLink}
      to='/'
      align='center'
      gap={2}
      _hover={{ textDecoration: 'none' }}
    >
      <Icon as={FaShoppingBag} w={iconSize} h={iconSize} color={color} />
      <Text
        fontSize={fontSize}
        fontWeight='bold'
        color={color}
        letterSpacing='tight'
      >
        Minh Plaza
      </Text>
    </Flex>
  );
};

export default Logo;
